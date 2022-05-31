" MapLeader

let mapleader = " "

" Plugins

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

	" Other
	Plug 'mattn/emmet-vim'
	Plug 'jiangmiao/auto-pairs'

	"Statusline
	Plug 'nvim-lualine/lualine.nvim'
	Plug 'kyazdani42/nvim-web-devicons'	

call plug#end()

" Configurations

set exrc
set hidden
set autoindent
set noexpandtab
set tabstop=2
set shiftwidth=2
set number relativenumber
set nowrap
set noswapfile
set nobackup
set undodir=~/.vim/undodir
set undofile
set incsearch
set scrolloff=8
set signcolumn=yes
set colorcolumn=120

" Color Schema

if (has("termguicolors"))
	set termguicolors
endif

syntax enable
colorscheme dracula

" Markdown Live Preview

let g:vim_markdown_folding_disabled=1
set nocompatible
filetype off

filetype plugin indent on
syntax enable
set nomore
set noswapfile
set viminfo=

let g:instant_markdown_logfile = '/tmp/instant_markdown.log'
"Uncomment to override defaults:
"let g:instant_markdown_slow = 1
"let g:instant_markdown_autostart = 0
"let g:instant_markdown_open_to_the_world = 1
"let g:instant_markdown_allow_unsafe_content = 1
"let g:instant_markdown_allow_external_content = 0
let g:instant_markdown_mathjax = 1
let g:instant_markdown_mermaid = 1
"let g:instant_markdown_autoscroll = 0
"let g:instant_markdown_port = 8888
"let g:instant_markdown_python = 1


" Vim CoC COnfiguration

source ~/.config/nvim/coc.vim

" LuaLine Config

source ~/.config/nvim/lualine.vim

" Mappings

inoremap <S-Tab> <C-d>
nnoremap <Leader>lor

" Telescope Mappings
nnoremap <Leader>ff <cmd>Telescope find_files<cr>
nnoremap <Leader>fg <cmd>Telescope live_grep<cr>
nnoremap <leader>fb <cmd>Telescope buffers<cr>
nnoremap <leader>fh <cmd>Telescope help_tags<cr>

" Custom Commands

function! GetFileTypeConfigPath() 
	let filetype = &filetype
	return "~/.config/nvim/ftplugin/" . &filetype . '.vim'
endfunction

command VimConfig e ~/.config/nvim/init.vim
command VimTypeConfig execute "e " . GetFileTypeConfigPath()
command SourceThis execute 'source ' . expand('%:p')

