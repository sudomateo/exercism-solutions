class Series

  private

  attr_reader :series

  public

  def initialize(series)
    @series = series
  end

  def slices(slice)
    raise ArgumentError if slice > series.chars.length
    series.chars.each_cons(slice).map(&:join)
  end
end
