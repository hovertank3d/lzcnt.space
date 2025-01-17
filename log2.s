    .section    .text
    .global     log2lzcnt
    .type       log2lzcnt, @function
log2lzcnt:
    lzcnt   %rdi, %rdi
    movq    $63,  %rax
    sub     %rdi, %rax
    ret
