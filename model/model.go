package model

type Command struct {
	Name   uint `gorm:"primaryKey"`
	SpecId uint `gorm:"primaryKey"`
}

type Spec struct {
	SpecId uint `gorm:"primaryKey"`

	// Spec HasMany(SpecId) Command
	Commands []Command `gorm:"foreignKey:SpecId;references:SpecId"`
}

type CommandResult struct {
	Name      uint `gorm:"primaryKey"`
	SpecId    uint `gorm:"primaryKey"`
	ContextId uint `gorm:"primaryKey"`

	// CommandResult BelongsTo(Name, SpecId) Command
	Command Command `gorm:"foreignKey:Name,SpecId"`
}

type Context struct {
	ContextId uint `gorm:"primaryKey"`

	// Context BelongsTo(SpecId) Spec
	SpecId uint `gorm:"primaryKey"`
	Spec   Spec

	// Context HasMany(ContextId, SpecId) CommandResult
	CommandResults []CommandResult `gorm:"foreignKey:ContextId,SpecId;references:ContextId,SpecId"`
}
