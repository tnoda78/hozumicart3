require 'RMagick'
require 'yaml'

class CartGifGenerater

  def initialize(letter, color, base_gif_file_path, position_file_path, font_file_path)
    @letter = letter
    @color = color
    @base_gif_file_path = base_gif_file_path
    @position_file_path = position_file_path
    @font_file_path = font_file_path
  end

  def gif_picture
    letter_one = @letter.length > 0 ? @letter[0] : ""
    letter_two = @letter.length > 1 ? @letter[1] : ""
    letter_three = @letter.length > 2 ? @letter[2] : ""

    position = YAML.load_file @position_file_path

    gif = Magick::ImageList.new @base_gif_file_path
    gif.each_with_index do |f, i|
      change_cart_color(f, @color)
      write_letter_to_image(f, i, letter_one, position["letter_one"], @font_file_path)
      write_letter_to_image(f, i, letter_two, position["letter_two"], @font_file_path)
      write_letter_to_image(f, i, letter_three, position["letter_three"], @font_file_path)
    end
    gif.to_blob
  end

  private

  def write_letter_to_image(img, i, letter, position, font_file_path)
    return if !(position[i]["enable"]) || letter == ""

    Magick::Draw.new.annotate(img, 0, 0, position[i]["x"], position[i]["y"], letter) do
      self.font = font_file_path
      self.align = Magick::CenterAlign
      self.stroke = "transparent"
      self.pointsize = 52
      self.text_antialias = true
      self.kerning = 1
    end
  end

  def change_cart_color(img, color)
    return if color == ""

    for y in 118...134
      for x in 0...img.columns
        src = img.pixel_color(x, y)
        r = src.red
        g = src.green
        b = src.blue
        if r == 37522 && g == 53456 && b == 20560
          img.pixel_color(x, y, "#" + color)
        end
      end
    end
  end
end

