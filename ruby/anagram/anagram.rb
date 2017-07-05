class BookKeeping
  VERSION = 2
end

class Anagram
  def initialize(word)
    @word = word.downcase
    @word_alpha_key = make_alpha_key(word)
  end

  def match(candidates)
    candidates.select { |candidate| anagram?(candidate) && not_the_word?(candidate)}
  end

  private

    def make_alpha_key(word)
      word.downcase.chars.sort.join
    end

    def anagram?(candidate)
      make_alpha_key(candidate) == @word_alpha_key
    end

    def not_the_word?(candidate)
      candidate.downcase != @word
    end
end
