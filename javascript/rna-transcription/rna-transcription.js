var DnaTranscriber = function(){
};

DnaTranscriber.prototype.toRna = function(dna){
  rna = '';
  for (var i=0; i<dna.length; i++){
    switch (dna[i]) {
      case 'C':
        rna += 'G';
        break;
      case 'G':
        rna += 'C';
        break;
      case 'A':
        rna += 'U';
        break;
      case 'T':
        rna += 'A';
        break;
    }
  }
  return rna;
};

module.exports = DnaTranscriber;
