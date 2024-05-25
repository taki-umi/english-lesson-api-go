package model

// User - ユーザー情報を定義します。
type User struct {
    userNumber      string
    name            string
    lessonCourse    string
    lessonStatus    string
    email           string
    lessonStartDate string
    baseEntity    BaseEntity
}

// NewUser - ユーザー情報を生成します。
func NewUser(
    userNumber string,
    name string,
    lessonCourse string,
    lessonStatus string,
    email string,
    lessonStartDate string,
    userID string,
) (*User, error) {

    be := NewBaseEntity(userID, userID)
    u := User{
        userNumber:      userNumber,
        name:            name,
        lessonCourse:    lessonCourse,
        lessonStatus:    lessonStatus,
        email:           email,
        lessonStartDate: lessonStartDate,
        baseEntity:    be,
    }

    err := u.validate()
    if err != nil {
        return nil, err
    }
    return &u, nil
}

// NewUserForRebuild - ユーザー情報を再構築します。
func NewUserForRebuild(
    userNumber string,
    name string,
    lessonCourse string,
    lessonStatus string,
    email string,
    lessonStartDate string,
    base BaseEntity,
) (*User, error) {

    u := User{
        userNumber:      userNumber,
        name:            name,
        lessonCourse:    lessonCourse,
        lessonStatus:    lessonStatus,
        email:           email,
        lessonStartDate: lessonStartDate,
        baseEntity:    base,
    }

    err := u.validate()
    if err != nil {
        return nil, err
    }
    return &u, nil
}

// validate - ユーザー情報のバリデーションを行います。
func (u *User) validate() error {
    return nil
}

// GetUserNumber - ユーザー番号を取得します。
func (u *User) UserNumber() string {
    return u.userNumber
}

// GetName - 名前を取得します。
func (u *User) Name() string {
    return u.name
}

// GetLessonCourse - レッスンコースを取得します。
func (u *User) LessonCourse() string {
    return u.lessonCourse
}

// GetLessonStatus - レッスンステータスを取得します。
func (u *User) LessonStatus() string {
    return u.lessonStatus
}

// GetEmail - メールアドレスを取得します。
func (u *User) Email() string {
    return u.email
}

// GetLessonStartDate - レッスン開始日を取得します。
func (u *User) LessonStartDate() string {
    return u.lessonStartDate
}

// BaseEntity - ベースエンティティを定義します。
func (u *User) BaseEntity() BaseEntity {
    return u.baseEntity
}

// ChangeName - 名前を変更します。
func (u *User) ChangeName(name string) error {
    current := u.name
    u.name = name

    err := u.validate()
    if err != nil {
        u.name = current
        return err
    }
    return nil
}

// ChangeLessonCourse - レッスンコースを変更します。
func (u *User) ChangeLessonCourse(lessonCourse string) error {
    current := u.lessonCourse
    u.lessonCourse = lessonCourse

    err := u.validate()
    if err != nil {
        u.lessonCourse = current
        return err
    }
    return nil
}

// ChangeLessonStatus - レッスンステータスを変更します。
func (u *User) ChangeLessonStatus(lessonStatus string) error {
    current := u.lessonStatus
    u.lessonStatus = lessonStatus

    err := u.validate()
    if err != nil {
        u.lessonStatus = current
        return err
    }
    return nil
}

// ChangeEmail - メールアドレスを変更します。
func (u *User) ChangeEmail(email string) error {
    current := u.email
    u.email = email

    err := u.validate()
    if err != nil {
        u.email = current
        return err
    }
    return nil
}

// ChangeLessonStartDate - レッスン開始日を変更します。
func (u *User) ChangeLessonStartDate(lessonStartDate string) error {
    current := u.lessonStartDate
    u.lessonStartDate = lessonStartDate

    err := u.validate()
    if err != nil {
        u.lessonStartDate = current
        return err
    }
    return nil
}
