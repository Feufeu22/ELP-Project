module Main exposing (..)
import Browser
import Html exposing (Html, button, div, input, text)
import Html.Events exposing (onClick, onInput)

-- MODEL

type alias Model =
    { title: String
    , inputText: String
    }

init : Model
init =
    { title = "Trouve la définition !"
    , inputText = ""
    }

-- UPDATE

type Msg
    = InputChanged String
    | DefinitionFound

update : Msg -> Model -> Model
update msg model =
    case msg of
        InputChanged newText ->
            { model | inputText = newText }
        DefinitionFound ->
            { model | title = "Vacances" }

-- VIEW

view : Model -> Html Msg
view model =
    div []
        [ text model.title        , div [] [ text "une période pendant laquelle une personne cesse ses activités habituelles." ]
        , input [ onInput InputChanged ] []
        , button [ onClick DefinitionFound ] [ text "Montrer définition" ]
        ]

main =
    Browser.sandbox { init = init, update = update, view = view }