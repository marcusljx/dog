package dog

type Setter interface {
	Set(key string, value interface{})
}

type Getter interface {
	Get(key string) (interface{}, error)
}

type GetSetter interface {
	Getter
	Setter
}

type Opener interface {
	Open(path string) error
}

type Closer interface {
	Close() error
}

type OpenCloser interface {
	Opener
	Closer
}
