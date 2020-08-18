// Simple Closure and Currying

function foo (outerMessage) {
    return function(innerMessage) { console.log(`Outer: ${outerMessage} \nInner: ${innerMessage}`) }
}

foo('Out')('And In !');
