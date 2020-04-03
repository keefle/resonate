package fuse

type Option func(*FS)

func GeneralOption(operation HookType, h GeneralHook) Option {
	return func(fuse *FS) {
		fuse.hooks[operation] = h
	}
}
