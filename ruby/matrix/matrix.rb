class Matrix
  def initialize(matrix)
    @matrix = matrix
  end

  def rows
    @matrix.split("\n").map do |row|
      row.split(" ").map do |elem|
        elem.to_i
      end
    end
  end

  def columns
    rows.transpose
  end
end
