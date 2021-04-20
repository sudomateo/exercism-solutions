class Phrase
  def initialize(phrase)
    @phrase = phrase
  end

  attr_reader :phrase

  def words
    phrase.downcase.scan(/\b[\w']+\b/)
  end

  def word_count()
    words.each_with_object(Hash.new(0)) do |word, count|
      count[word] += 1
    end
  end
end
