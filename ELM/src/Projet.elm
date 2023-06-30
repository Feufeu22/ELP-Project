-------------------------
-- *  FIND THE WORD  * --
-------------------------

module Projet exposing (..)
import Browser
import Http exposing(..)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Random
import List exposing (..)
import String exposing (words)
import Json.Decode exposing (..)
import Platform.Cmd exposing (..)
import Browser.Dom exposing (..)
import Json.Decode.Pipeline as JP

--------------------------------------------------------------------------------------------------------
-- MAIN

main =
  Browser.element { init = init, update = update, subscriptions = subscriptions, view = view }

--------------------------------------------------------------------------------------------------------
-- MODEL

type  Model 
  = Loading
  | Success Data

type alias Data 
  = { text : String
    , definition : Def
    , inputUser : String
    , color : String    
    , tempInput : String
    , reponse : String
    , showResponse : String
    , state : String
    , resultat : String
    , essai : Int
    }
  
--------------------------------------------------------------------------------------------------------
-- HTTP

init : ()->(Model, Cmd Msg)
init _ =
   (Loading
   , Http.get{ url = "http://localhost:8000/mots.txt", expect = Http.expectString GotText}
  )

getDef : String -> Cmd Msg
getDef mot =
  Http.get
    { url = "https://api.dictionaryapi.dev/api/v2/entries/en/"++mot
    , expect = Http.expectJson GotDef (Json.Decode.list defDecoder)
    }

--------------------------------------------------------------------------------------------------------
-- JSON

type alias Meaning =
    { partOfSpeech : String
    , definitions : List Definition
    }

type alias Definition =
    { definition : String
    , synonyms : List String
    , antonyms : List String
    }

type alias Def =
    { meanings : List Meaning
    }

type alias API =
    {def : List Def}

defMeaDecoder =
    succeed Definition
        |> JP.required "definition" string
        |> JP.required "synonyms" (Json.Decode.list string)
        |> JP.required "antonyms" (Json.Decode.list string)

defDecoder =
    succeed Def
        |> JP.required "meanings" (Json.Decode.list meaningDecoder)

meaningDecoder =
    succeed Meaning
        |> JP.required "partOfSpeech" string
        |> JP.required "definitions" (Json.Decode.list defMeaDecoder)

--------------------------------------------------------------------------------------------------------
-- UPDATE

type Msg
  = InputUser String
  | GotText (Result Http.Error String)
  | GotDef (Result Http.Error (List Def))
  | Change
  | Test
  | Random Int
  | Show

rool : Model -> Cmd Msg
rool model =
  case model of
    Loading ->
      rool model
    Success data ->
      Random.generate Random (Random.int 0 (length(words data.text)))

randomWord : Int -> Model -> String
randomWord index model =
  case model of
      Loading ->
        " "
      Success data ->
        let mot = head (drop index (words data.text)) in case mot of
          Just str -> str
          Nothing -> "ERROR"


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case model of
    Loading ->
      case msg of 
      GotText result ->
        case result of
          Ok fullText ->
             update Change (Success{ text = fullText
                                    , definition = {meanings = []}
                                    , inputUser =""
                                    , tempInput =""
                                    , reponse =""
                                    , essai =0
                                    , state ="GotText"
                                    , resultat =""
                                    , color =""
                                    ,showResponse = ""})
          Err _->
            (model,Cmd.none)
      _ -> (model,Cmd.none)

    Success data ->  
      case msg of 
        InputUser temp ->
          (Success{ data | inputUser = temp, state = "input" },Cmd.none)

        Change ->
          (Success{data | essai = 0
          , resultat = " "
          , inputUser = ""
          , state = "change"
          ,showResponse = ""
          }, rool model)

        Random result ->
          (Success{data | reponse = (randomWord result model)
            }, getDef (randomWord result model))

        Test ->
          (Success{data | resultat = 
          if data.reponse == data.inputUser then
            "Bonne Réponse"  
          else
            "Bad answer"
          ,color = 
          if data.reponse == data.inputUser then 
            "green"  
          else
            "red"
          ,essai = 
          if data.inputUser /= data.reponse then 
            data.essai +1
          else
            data.essai
          , state = 
          if data.reponse == data.inputUser then
            "testTrue"
          else
            "testFalse"}, Cmd.none)

        Show ->
          (Success{data | showResponse = ("   : "++data.reponse)
          , state = "testTrue"
          }, Cmd.none)  

        GotDef result -> 
          case result of
              Ok def ->
                case (head def) of
                  Nothing -> (model, Cmd.none)
                  Just rs -> (Success{data | definition = rs}, Cmd.none)
              Err _ ->
                (model, Cmd.none)
        
        GotText result ->
          case result of
            Ok fullText ->
              (Success{ data | text = fullText}, rool model)
            Err _->
              (model,Cmd.none)         

--------------------------------------------------------------------------------------------------------
-- VIEW

view : Model -> Html Msg
view model =
  div[]
   [ h1[] [text "Find the word"]
   , viewPage model
   ]

viewPage : Model -> Html Msg
viewPage model =
  case model of
    Loading -> div[][viewText ("Loading..... ") "h1"]
    Success data ->
      div[]
        [ viewTry model
        , br[][]
        , viewDef data.definition
        , div[][viewText  "Your response here" "h2"
                  ,viewInput "text"  data.inputUser InputUser
                  ,viewBouton "Test" Test         
            ]
        , br[][]
        , div[][viewValidation model, viewBouton "Show answer" Show, viewText (data.showResponse) ""
              ]   
        ]

viewTry : Model -> Html Msg
viewTry model =
  case model of
    Loading -> div[][]
    Success data ->
      div[]
      [ viewText "See definition here" "h2"
      ]

viewDef : Def -> Html Msg
viewDef def  =
    ul[][
      li[][ text "meaning", br[][] ,
        ul[][
          div[](List.map viewMeanings def.meanings)
        ]
      ]
    ]

viewUnderDef : Definition -> Html Msg
viewUnderDef underDef =
    li[][
        viewText underDef.definition ""
    ]


viewListString : List String -> Html Msg
viewListString listString =
    div[](List.map text listString)


viewMeanings : Meaning -> Html Msg
viewMeanings  meanings =
  li[][ viewText meanings.partOfSpeech "h4",
    div[][ol[](List.map viewUnderDef meanings.definitions)]
  ]


viewInput : String -> String -> (String -> msg) -> Html msg
viewInput t v toMsg =
  input [ type_ t, Html.Attributes.value v, onInput toMsg ] []


viewText : String -> String ->  Html msg
viewText t style =
  case style of 
    "h1"->
      h1[] [text t]
    "h2"->
      h2[] [text t]
    "h3"->
      h3[] [text t]
    "h4"->
      h4[] [text t]
    "h5"->
      h5[] [text t]
    "h6"->
      h6[] [text t]
    _ -> 
      text t
   

viewBouton : String -> (msg) -> Html msg
viewBouton t toMsg =
  button [ onClick toMsg ] [ text t ]


viewValidation : Model -> Html msg
viewValidation model =
  case model of
    Loading -> div[][]
    Success data ->
      div [ style "color" data.color] [ text data.resultat ]


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none