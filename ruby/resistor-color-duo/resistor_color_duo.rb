class ResistorColorDuo
  COLORS = ["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"]

  def self.value(bands)
    if bands.length < 2
      return 0
    end
    COLORS.find_index(bands[0]) * 10 + COLORS.find_index(bands[1])
  end
end
