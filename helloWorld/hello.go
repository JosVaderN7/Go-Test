package hello

//Le falta agregar 2 parametros
//Una variable String con un nombre
func Hello(name string, lang string) string {
	
	if name == "" {
		return "Hello, World!"
	} 
	

	if lang == "Spanish"{
	 return "Hola, Jorch!"
	}

	if lang == "French"{
		return "Bonjour, Jorch!"
	}
	return "Hello, " + name +"!"
}
