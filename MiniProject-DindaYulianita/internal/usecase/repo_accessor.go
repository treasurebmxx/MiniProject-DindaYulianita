package usecase
import "MiniProject-DindaYulianita/internal/repository"


func (u *Usecase) Repo() *repository.Repository {
    return u.repo
}
