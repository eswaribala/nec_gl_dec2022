package facades

// IEcommFacade interface
type IECommFacade interface {
	Create()
	//View abstract method
	View(permission bool)
}
