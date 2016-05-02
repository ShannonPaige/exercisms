var Gigasecond = function(oldDate){
  this.oldDate= oldDate
};

Gigasecond.prototype.date = function() {
  var newDay = new Date();
  newDay.setTime(this.oldDate.getTime()+1000000000000);
  return newDay
}


module.exports = Gigasecond;
