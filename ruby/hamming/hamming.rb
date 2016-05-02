class Hamming
  VERSION = 1

  def self.compute(string1, string2)
    fail ArgumentError if string1.length != string2.length
    hamming_number = 0
    string1.chars.each_with_index do |letter, i|
      letter == string2[i] ? hamming_number : hamming_number+=1
    end
    hamming_number
  end                                          # => :compute
end                                            # => :compute
