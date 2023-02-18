// This the the toy project to demonstrate how to write a testable program.
// chao.cai@mobvista.com
package toy

// Storage is to persist the data
type Storage interface {
	// RetiveData is to retrive the data by the associated token.
	RetiveData(token string) (string, error)

	// StoreData is to persist the data
	StoreData(token string, data string) error
}

// Processor is to provide the processing service
type Processor interface {
	// Process is to the raw data
	Process(raw string) (string, error)
}

type TokenCreator interface {
	CreateToken(data string) string
}

type ProcessingService struct {
	Storage      Storage
	Processor    Processor
	TokenCreator TokenCreator
}

func NewProcessingService(processor Processor,
	tokenCreator TokenCreator, storage Storage) *ProcessingService {
	return &ProcessingService{
		Storage:      storage,
		Processor:    processor,
		TokenCreator: tokenCreator,
	}
}

func (ps *ProcessingService) Process(raw string) (string, error) {
	processed, err := ps.Processor.Process(raw)
	if err != nil {
		return "", err
	}
	token := ps.TokenCreator.CreateToken(processed)
	err1 := ps.Storage.StoreData(token, processed)
	return token, err1
}

func (ps *ProcessingService) Retrive(token string) (string, error) {
	return ps.Storage.RetiveData(token)
}
