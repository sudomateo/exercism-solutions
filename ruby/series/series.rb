class Series

  attr_reader :series

  def initialize(series)
    @series = series
  end

  def slices(slice)
    raise ArgumentError if slice > @series.chars.length
    @series.chars.each_cons(slice).map(&:join)
  end
end
