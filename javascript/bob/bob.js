var Bob = function() {};

Bob.prototype.hey = function(input) {
  var input = input.replace(/\s+/g, '')

  if (input === '') {
    return 'Fine. Be that way!'
  } else if (input === input.toUpperCase() && input !== input.toLowerCase()){
    return 'Whoa, chill out!'
  } else if (input.charAt(input.length-1) === '?'){
    return 'Sure.'
  } else {
    return 'Whatever.'
  }
};

module.exports = Bob;
