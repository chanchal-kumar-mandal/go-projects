# Create Custom Package

* Details: We will structure the code in such a way that all functionalities related to a rectangle are in rectangle package.

    Lets create a custom package rectangle which has functions to determine the area and diagonal of a rectangle.

    Source files belonging to a package should be placed in separate folders of their own. It is a convention in Go to name this folder       with the same name of the package.

    So lets create a folder named rectangle inside the geometry folder. All files inside the rectangle folder should start with the line       package rectangle as they all belong to the rectangle package.

    Create a file rectprops.go inside the rectangle folder we just created and add the following code.

* Folder Structure:    
    
    src
        
        geometry
            
            geometry.go
            
            rectangle
                
                rectprops.go
                     
* How To Run?
    1. Go to path /C/Users/[chanchal]/Go/src/geometry in git bash
    2. Run Command: $ go install geometry
    3. Run Command: $ go run geometry.go
    
* Source:
    https://golangbot.com/packages/


                     
              
