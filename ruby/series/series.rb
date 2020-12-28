class Series
  def initialize(series)
    @series = series.chars
  end

  def slices(slice)
    raise ArgumentError if slice > @series.length
    @series.each_cons(slice).map { |chars| chars.join }
  end
end
