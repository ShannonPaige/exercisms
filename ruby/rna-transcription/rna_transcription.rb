class Complement
  VERSION = 2

  def self.of_dna(dna)
    fail ArgumentError if (/[^GCTA]/.match(dna))

    rna = dna.gsub("G", "0").gsub("C", "1").gsub("T", "2").gsub("A", "3")
    rna = rna.gsub("0", "C").gsub("1", "G").gsub("2", "A").gsub("3", "U")
    return rna
  end

  def self.of_rna(rna)
    fail ArgumentError if (/[^CGAU]/.match(rna))

    dna = rna.gsub("C", "0").gsub("G", "1").gsub("A", "2").gsub("U", "3")
    dna = dna.gsub("0", "G").gsub("1", "C").gsub("2", "T").gsub("3", "A")
    return dna
  end
end
