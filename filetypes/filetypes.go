package filetypes

type Export struct {
    Type string `json:"type"`
    Body string `json:"body"`
}

type Logfile struct {
    Notes []string `json:"notes"`
    Replacements map[string]string `json:"replacements"`
    Missions []Mission `json:"missions"`
}

type Mission struct {
    Date string `json:"date"`
    Time string `json:"time"`
    Keyword string `json:"keyword"`
    Situation string `json:"situation"`
    Unit string `json:"unit"`
    Location string `json:"location"`
    Links []string `json:"links"`
}
