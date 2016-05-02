var Year = function(year) {
  this.year = year;
};

Year.prototype.isLeap = function(year) {
  if (this.year % 4 === 0 && this.year % 100 !== 0 || this.year % 400 === 0) {
    return true
  } else {
    return false
  }
};

module.exports = Year;
