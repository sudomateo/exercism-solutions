class Acronym
  def self.abbreviate(str)
    acronym = ""
    str.split(" ").map { |word|
      if word.include? "-"
        word.split("-").map { |sub_word|
          acronym += sub_word[0].upcase
        }
      else
        acronym += word[0].upcase
      end
    }
    acronym
  end
end
