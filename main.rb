require 'sinatra'
require 'sinatra/base'

class MainApp < Sinatra::Base
  get '/' do
    'Hello, World! And Hello.And Hello.'
  end
end
