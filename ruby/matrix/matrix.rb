class Matrix
  def initialize(matrix)
    @matrix = matrix
  end

  def rows
    @matrix.each_line.map do |row|
      row.split().map do |elem|
        elem.to_i
      end
    end
  end

  def columns
    rows.transpose
  end
end
