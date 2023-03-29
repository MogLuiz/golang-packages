package main

import "net/http"

// O Mux serve para criar rotas e servidores customizáveis e com um maior controle sobre as rotas do nosso sistema. Uma vez que usando a forma padrão de levantar rota
// como no exemplo da pasta "6-http-server" aquela é uma instância global do Go e do projeto.
// Com o MUX eu posso criar diversos servidores, diversas portas o que não é possível na forma padrão: ex: Porta 8080 com a rota /Home, porta 6000 com a rota /Home etc...
// Pois na forma padrão estamos utilizando o mesmo serve MUX.
// Aqui, podemos instanciar vários server's MUX

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Blog customizável"})
	http.ListenAndServe(":8080", mux)

	// Criando um segundo servidor
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Olá segundo server mux!"))
	})
	http.ListenAndServe(":8081", mux2)

}

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Olá mundo!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(b.title))
}
