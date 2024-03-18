class Task_manager
    attr_accessor :tasks, :next_id

    def initialize()
        @tasks = {}
        @next_id = 1
    end

    def loop()
        while true do
            puts "Enter a command or type `help`"
            input = gets.chomp
            case input
            when "add"
                create_task()
            when "view"
                display_tasks
            when "help"
                puts "The available commands are:\nadd\nview\nhelp\ndelete\nexit"
            when "delete"
                delete_task()
            when "exit"
                puts "Here is your final Todo List:"
                display_tasks()
                puts "Goodbye!"
                break 
            end
            input = ""
        end
    end

    def delete_task()
        puts "Which task would you like to delete?"
        display_tasks()
        input = gets.chomp
        begin
            input = input.to_i
            @tasks.delete(input)
            puts "Task deleted successfully!"
        rescue
            puts "Invalid input. No tasks deleted."
        end
    end

    def create_task()
        puts "Please enter the text for your task:"
        input = gets.chomp
        @tasks[@next_id] = input
        @next_id += 1
    end

    def display_tasks()
        for id, task in @tasks
            puts "#{id}.) #{task}"
        end
        puts
    end
end
