package expr

import (
    "fmt"
    "github.com/ThCompiler/go_game_constractor/pkg/cleanenv"
    "github.com/ThCompiler/go_game_constractor/pkg/graph"
    "github.com/ThCompiler/go_game_constractor/scg/expr/scene"
    "github.com/ThCompiler/go_game_constractor/scg/script/matchers"
    "github.com/pkg/errors"
)

type Script map[string]scene.Scene

type ScriptInfo struct {
    StartScene     string                         `yaml:"startScene" json:"start_scene" xml:"startScene"`
    Name           string                         `yaml:"name" json:"name" xml:"name"`
    GoodByeCommand string                         `yaml:"goodByeCommand" json:"good_bye_command" xml:"goodByeCommand"`
    GoodByeScene   string                         `yaml:"goodByeScene" json:"good_bye_scene" xml:"goodByeScene"`
    UserMatchers   map[string]scene.ScriptMatcher `yaml:"matchers,omitempty" json:"matchers,omitempty" xml:"matchers,omitempty"`
    Script         Script                         `yaml:"script" json:"script" xml:"script"`
}

func (si *ScriptInfo) IsValid() (is bool, err error) {
    if is, err = si.checkValidScenes(); !is {
        return is, err
    }

    if is, err = si.checkNextScenes(); !is {
        return is, err
    }

    if is, err = si.checkValidMatcherName(); !is {
        return is, err
    }

    if _, is = si.Script[si.GoodByeScene]; !is {
        return false, ErrorGoodbyeSceneNotFound
    }

    if _, is = si.Script[si.StartScene]; !is {
        return false, ErrorStartSceneNotFound
    }

    return false, nil
}

func (si *ScriptInfo) checkValidScenes() (is bool, err error) {
    for _, sc := range si.Script {
        if is, err = sc.IsValid(si.UserMatchers); !is {
            break
        }
    }
    return is, err
}

func (si *ScriptInfo) checkValidMatcherName() (is bool, err error) {
    for _, matcher := range si.UserMatchers {
        name := ""
        if matcher.IsRegexMatcher() {
            name = matcher.MustRegexMatcher().Name
        } else {
            name = matcher.MustSelectsMatcher().Name
        }

        if is = matchers.IsCorrectNameOfMather(name); !is {
            err = errors.Wrap(ErrorNameAlreadyOccupied, fmt.Sprintf("error with matcher %s", name))
            break
        }
    }

    return is, err
}

func (si *ScriptInfo) checkNextScenes() (is bool, err error) {
    unknownScene := ""
up:
    for _, sc := range si.Script {

        if _, is = si.Script[sc.NextScene]; sc.IsInfoScene && !is {
            unknownScene = sc.NextScene
            break
        }

        for _, name := range sc.NextScenes {
            if _, is = si.Script[name]; !is {
                unknownScene = name
                break up
            }
        }
    }

    if !is {
        if unknownScene == "" {
            return is, err
        }
        return is, errorNameSceneNotFound(unknownScene)
    }

    return false, nil
}

type sceneContext struct {
    ctx       scene.Context
    sceneName string
}

func (si *ScriptInfo) checkContext() (bool, error) {
    sceneGraph, err := si.initSceneGraph()
    if err != nil {
        return false, err
    }

    if is, err := si.checkLoadContext(sceneGraph); !is {
        return is, err
    }

    if is, err := si.checkValueContext(sceneGraph); !is {
        return is, err
    }

    return err == nil, err
}

func (si *ScriptInfo) checkLoadContext(sceneGraph *graph.Graph[*sceneContext, string]) (bool, error) {
    err := error(nil)
up:
    for name, sc := range si.Script {
        for _, load := range sc.Context.LoadValue {
            visited := make([]string, 0)
            found := false

            visitor := graph.Visitor[*sceneContext](func(sctx *graph.ValueNode[*sceneContext]) bool {
                visited = append(visited, sctx.Value.sceneName)
                if sctx.Value.ctx.SaveValue.Name == load.Name {
                    found = true
                    load.Type = sctx.Value.ctx.SaveValue.Type
                    return true
                }
                return false
            })

            sceneGraph.BFS(name, visitor)

            if !found {
                err = errorNotFoundLoadingContext(load.Name, name, visited)
                break up
            }
        }
    }

    return err == nil, err
}

func (si *ScriptInfo) checkValueContext(sceneGraph *graph.Graph[*sceneContext, string]) (bool, error) {
    err := error(nil)
up:
    for name, sc := range si.Script {
        for _, value := range sc.Text.Values {
            visited := make([]string, 0)
            found := false
            ctxType := ""

            visitor := graph.Visitor[*sceneContext](func(sctx *graph.ValueNode[*sceneContext]) bool {
                visited = append(visited, sctx.Value.sceneName)
                if sctx.Value.ctx.SaveValue.Name == value.FromContext {
                    found = true
                    ctxType = sctx.Value.ctx.SaveValue.Type
                    return true
                }
                return false
            })

            sceneGraph.BFS(name, visitor)

            if !found {
                err = errorNotFoundLoadingContextInValues(value.FromContext, name, visited)
                break up
            }

            if ctxType != value.Type {
                err = nil
                break up
            }
        }
    }

    return err == nil, err
}

func (si *ScriptInfo) initSceneGraph() (*graph.Graph[*sceneContext, string], error) {
    sceneGraph := graph.NewGraph[*sceneContext, string](nil)
    edges := make([]graph.EdgeInfo[string], 0)

    for name, sc := range si.Script {
        sceneGraph.AddVertex(name, &sceneContext{sc.Context, name})

        for _, nextScene := range sc.NextScenes {
            edges = append(edges, graph.EdgeInfo[string]{VertexFrom: nextScene, VertexTo: name})
        }

        if sc.NextScene != "" {
            edges = append(edges, graph.EdgeInfo[string]{VertexFrom: sc.NextScene, VertexTo: name})
        }
    }

    err := sceneGraph.AddUndirectedEdges(edges...)
    if err != nil {
        return nil, errors.Wrap(ErrorUnknown, err.Error())
    }
    return sceneGraph, nil
}

func LoadScriptInfo(path string) (*ScriptInfo, error) {
    si := ScriptInfo{}
    err := cleanenv.ReadConfig(path, &si)
    if err != nil {
        return nil, fmt.Errorf("error load script info: %w", err)
    }

    _, err = si.IsValid()
    if err != nil {
        return nil, fmt.Errorf("this script config is not correct: %w", err)
    }
    return &si, nil
}
