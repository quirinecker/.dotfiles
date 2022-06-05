function! GetFileTypeConfigPath() 
	let filetype = &filetype
	return "~/.config/nvim/ftplugin/" . &filetype . '.vim'
endfunction

command VimConfig e ~/.config/nvim/init.vim
command VimTypeConfig execute "e " . GetFileTypeConfigPath()
