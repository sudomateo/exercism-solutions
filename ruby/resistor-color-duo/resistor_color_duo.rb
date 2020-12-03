class ResistorColorDuo
  @@resistance = {
    "black" => 0,
    "brown" => 1,
    "red" => 2,
    "orange" => 3,
    "yellow" => 4,
    "green" => 5,
    "blue" => 6,
    "violet" => 7,
    "grey" => 8,
    "white" => 9,
  }

  def self.value(bands)
    if bands.length < 2
      return 0
    end
    @@resistance[bands[0]] * 10 + @@resistance[bands[1]]
  end
end
