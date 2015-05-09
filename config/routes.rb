UwrTournament::Application.routes.draw do
  # You can have the root of your site routed with "root"
  root 'games#new'

  devise_for :users

  get '/statistics', to: 'public#statistics'

  resources :clubs, only: [:show, :new, :create]

  resources :games, only: [:create, :new, :show, :index]

  resources :scores, only: [:create]

  resources :teams, only: [] do
    get 'players', on: :member
  end
end
