    .section    .text
    .global     log2lzcnt
    .type       log2lzcnt, @function
# {{template "source_code.html" .}}
# {{template "reset.html" .}}
log2lzcnt:              # log2lzcnt({{.Requests}});
    lzcnt   %rdi, %rdi 
    movq    $63,  %rax
    sub     %rdi, %rax
    ret                 # 2^{{.Log2lzcnt}} requests handled
