# Changelog

# [1.2.0] - 2020-01-06

### Changed
   * New parameter `io.Reader` was added in method `Get` and `GetWithContext`.
     
        > Old:
            `func (f *Fetch) Get(url string) (*Response, error)`
        
        > New:
            `func (f *Fetch) Get(url string, reader io.Reader) (*Response, error)`
        
        > Old:
            `func (f *Fetch) GetWithContext(ctx context.Context, url string) (*Response, error)`
        
        > New:
            `func (f *Fetch) GetWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error)`
        
   * The method `response.String` was changed.
         
        > Old: 
             `func (r *Response) String() (string, error)`

        > New: 
             `func (r *Response) String() string`
    
   * New method ws added `response.ToString`
   
        > Method
            `func (r *Response) ToString() (string, error)` 