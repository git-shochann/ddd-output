package usecase

// 1.まず何のメソッドを定義して、直接usecase層に依存するところ(interface層)で使いたいか？
// 2.構造体にてこの層の責務を表現する

type UserUseCase interface {
}

type userUseCase struct {
}
