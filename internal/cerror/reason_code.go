package cerror

type ReasonCode string

const (
	RC00000 ReasonCode = "RC00000" // Message: reason code が設定されていないエラーです 対応策については実装を確認し、対応が必要であった場合は適切にコードを設定してください
	RC00001 ReasonCode = "RC00001" // Message: Panicエラーが発生しました 直ちにエンジニアは対応してください
	RC10100 ReasonCode = "RC10100"
	RC10101 ReasonCode = "RC10101"
	RC10200 ReasonCode = "RC10200"
)
