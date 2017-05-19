class Raindrops

  def self.convert(drops)
    sound = ""

    if drops % 3 == 0
      sound += "Pling"
    end
    if drops % 5 == 0
      sound += "Plang"
    end
    if drops % 7 == 0
      sound += "Plong"
    end

    if sound == ""
      sound = drops.to_s
    end

    return sound
  end
end

class BookKeeping
  VERSION = 3
end
