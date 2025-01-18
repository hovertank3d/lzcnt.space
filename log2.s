# {{template "link.html" (arr "about" "/about")}}
# {{template "reset.html" .}}
    .section    .text
    .global     log2lzcnt
    .type       log2lzcnt, @function
log2lzcnt:              # log2lzcnt({{.Requests}});
    lzcnt   %rdi, %rdi 
    movq    $63,  %rax
    sub     %rdi, %rax
    ret                 # 2^{{.Log2lzcnt}} requests handled

# copyleft 2025 iskrim46
# {{template "link.html" (arr "source" "https://github.com/hovertank3d/lzcnt.space")}}
