package goliblocale

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

func GetLocale (loc string) (map[string]string, error) {
  var trans map[string]string

  data, err := ioutil.ReadFile(fmt.Sprintf("%s.json", loc))
  if err != nil {
    if os.IsNotExist(err) {
      fmt.Printf("言語 '%s' を見つけられませんでした。 'ja' を利用します。\n", loc)
      data, err = ioutil.ReadFile("ja.json")
      if err != nil {
        return trans, err
      }
    } else {
      return trans, err
    }
  }

  err = json.Unmarshal(data, &trans)
  if err != nil {
    return trans, err
  }

  return trans, nil
}
