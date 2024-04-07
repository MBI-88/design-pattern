package cmd



type Adapter struct {
	admap   *adapteeToMap
	adarray *adapteeToArray
}


func (a *Adapter) SetParams(name, lastname string, id int) {
	a.admap = new(adapteeToMap)
	a.adarray = new(adapteeToArray)
	a.admap.SetParameters(name, lastname, id)
	a.adarray.SetParameters(name, lastname, id)
}

func (a Adapter) GetArrayParams() []string {
	return a.adarray.GeParams()
}

func (a Adapter) GetMapParams() map[string]string {
	return a.admap.GeParams()
}