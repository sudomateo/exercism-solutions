class TwoFer
  def self.two_fer(name = "")
    if name == ""
      return "One for you, one for me."
    end
    return "One for %s, one for me" % name
  end
end
