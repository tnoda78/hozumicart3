require 'sinatra'
require 'sinatra/base'
require 'erb'
require './lib/cart_gif_generater'

class MainApp < Sinatra::Base

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
      #get_gif_picture(letter, color)
      generater = CartGifGenerater.new(
        letter,
        color,
        "images/cart.gif",
        "config/position.yml",
        "fonts/ipagp.ttf"
      )
      generater.gif_picture
    end
  end
end
