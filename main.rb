require 'sinatra'
require 'sinatra/base'
require 'erb'
require 'RMagick'
require 'yaml'

class MainApp < Sinatra::Base

  helpers do
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

    def write_letter_to_image(img, i, letter, position)
      return if !(position[i]["enable"]) || letter == ""

      Magick::Draw.new.annotate(img, 0, 0, position[i]["x"], position[i]["y"], letter) do
        self.font = "fonts/ipagp.ttf"
        self.align = Magick::CenterAlign
        self.stroke = "transparent"
        self.pointsize = 52
        self.text_antialias = true
        self.kerning = 1
      end
    end

    def get_gif_picture(letter, color)
      letter_one = letter.length > 0 ? letter[0] : ""
      letter_two = letter.length > 1 ? letter[1] : ""
      letter_three = letter.length > 2 ? letter[2] : ""

      position = YAML.load_file "config/position.yml"

      gif = Magick::ImageList.new "images/cart.gif"
      gif.each_with_index do |f, i|
        change_cart_color(f, color)
        write_letter_to_image(f, i, letter_one, position["letter_one"])
        write_letter_to_image(f, i, letter_two, position["letter_two"])
        write_letter_to_image(f, i, letter_three, position["letter_three"])
      end
      gif.to_blob
    end
  end

  get '/' do
    @base_url = "#{request.scheme}://#{request.host}:#{request.port}/cart"
    erb :index
  end

  get '/cart' do
    letter = params[:letter] || ""
    color = params[:color] || ""
    content_type 'image/gif'
    if letter == "" && color == ""
      send_file "images/cart.gif"
    else
      get_gif_picture(letter, color)
    end
  end
end
