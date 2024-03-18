require_relative "task_manager"

def greet()
    puts "Welcome to the Ruby Todo List!"
end

greet()
m = Task_manager.new
m.loop()
