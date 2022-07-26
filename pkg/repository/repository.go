package repository

type Repository struct {
	GitUrl string
	Watch  bool
}

func (r *Repository) Clone() error {
	// TODO: Clone the repository.
	return nil
}
