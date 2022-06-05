
call plug#begin("~/.vim/plugged")

	" Basics
	Plug 'junegunn/vim-plug'

	" Markdown
	Plug 'instant-markdown/vim-instant-markdown', {'for': ['markdown', 'markdown.pandoc']}
	Plug 'godlygeek/tabular'
	Plug 'preservim/vim-markdown'

	" coc
	Plug 'neoclide/coc.nvim', {'branch': 'release'}

	" Telescope
	Plug 'nvim-telescope/telescope.nvim'
	Plug 'nvim-lua/plenary.nvim'
	Plug 'nvim-treesitter/nvim-treesitter', {'do': ':TSUpdate'}

	" Themes
	Plug 'dracula/vim'

	"Statusline
	Plug 'nvim-lualine/lualine.nvim'
	Plug 'kyazdani42/nvim-web-devicons'	

	" Other
	Plug 'mattn/emmet-vim'
	Plug 'jiangmiao/auto-pairs'
	Plug 'posva/vim-vue'
	Plug 'editorconfig/editorconfig-vim'
	Plug 'preservim/nerdtree'

call plug#end()
