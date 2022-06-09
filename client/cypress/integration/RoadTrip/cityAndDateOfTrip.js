describe('The cities and dates of the trip', () => {

    it('the start city is taken from the previous page', () => {
        cy.visit('http://localhost:3000') // change URL to match your dev URL
        cy.wait(500)
        cy.fixture('tripInformations').then((tripInformations) => {
            cy.get('#enterStartCitySearchBar').type(tripInformations.startCity)

            //SÉLECTIONNER LA PREMIÈRE VALEUR DANS LA LISTE DÉROULANTE
            cy.get('#first-choice').click()
            cy.wait(2000)
            //check if it's the good page
            cy.location('pathname').should('include', '/cities')

            //check the input of the start city
            cy.get('#startCity').contains(tripInformations.startCity)
        })
    })





    it('check if the next button is disable if all information is not present enter', () => {
        cy.visit('http://localhost:3000/cities') // change URL to match your dev URL
        cy.wait(500)

        cy.get('#nextButton').should('be.disabled')

        cy.fixture('tripInformations').then((tripInformations) => {
            //Start city
            //clear beacause this field is already filled in normally
            cy.get('#startCity').clear()
            cy.get('#startCity').type(tripInformations.startCity)
            cy.get('#saveButton').should('be.disabled')

            //End city
            cy.get('#endCity').type(tripInformations.endCity)
            cy.get('#saveButton').should('be.disabled')

            //Date
            cy.get('#startDate').type(tripInformations.startDate)
            cy.get('#saveButton').should('be.disabled')
            cy.get('#EndDate').type(tripInformations.endDate)


            cy.get('#saveButton').should('not.be.disabled') //all information are here

            cy.wait(1000)
            cy.get('#saveButton').click()
            cy.location('pathname').should('include', '/travels')
        })
    })



    it('select travel', () => {
        cy.visit('http://localhost:3000/travels') // change URL to match your dev URL
        cy.wait(500)

        //check button return
        cy.get('#returnButton').click()
        cy.wait(1000)
        cy.location('pathname').should('include', '/cities')
        cy.get('#saveButton').click()
        cy.wait(1000)
        cy.location('pathname').should('include', '/travels')

        //check Save button is disable because we don't choice travel
        cy.get('#saveButton').should('be.disabled')

        //check different travels button
        cy.get('#Marche').click()
        cy.wait(500)
        cy.get('#Vélo').click()
        cy.wait(500)
        cy.get('#Voiture').click()
        cy.wait(500)
        cy.get('#Avion').click()
        cy.wait(500)

        //Select a travel and Save
        cy.get('#travel1').click()
        cy.wait(500)

        cy.get('#saveButton').should('not.be.disabled')
        cy.wait(500)
        cy.get('#saveButton').click()
        cy.wait(1000)

        cy.location('pathname').should('include', '/activities')
    })





    it('select activities Start city', () => {
        cy.visit('http://localhost:3000/activities/Start') // change URL to match your dev URL
        cy.wait(500)

        //user can go to the next page without save an activities
        cy.get('SaveButton').should('not.be.disable')

        //check button return
        cy.get('#returnButton').click()
        cy.wait(1000)
        cy.location('pathname').should('include', '/travels')
        cy.get('#saveButton').click()
        cy.wait(1000)
        cy.location('pathname').should('include', '/activities/Start')

        //check the name of the city on this page
        cy.fixture('tripInformations').then((tripInformations) => {
            cy.get('#HeaderPage').contains(tripInformations.startCity)
        })

        //check different activities button
        cy.get('#Hôtels').click()
        cy.wait(500)
        cy.get('#LocationVoitures').click()
        cy.wait(500)
        cy.get('#Activités').click()
        cy.wait(500)
        cy.get('#Restaurants').click()
        cy.wait(500)
        cy.get('#Croisières').click()
        cy.wait(500)

        //Add an activity for each type of activity


        //Delete one activities


        //go to the next page
        cy.get('SaveButton').click()
    })




    it('select activities End city', () => {
        cy.visit('http://localhost:3000/activities/End') // change URL to match your dev URL
        cy.wait(500)

        //user can go to the next page without save an activities
        cy.get('SaveButton').should('not.be.disable')

        //check button return
        cy.get('#returnButton').click()
        cy.wait(1000)
        cy.location('pathname').should('include', '/activities/Start')
        cy.get('#saveButton').click()
        cy.wait(1000)
        cy.location('pathname').should('include', '/activities/End')

        //check the name of the city on this page
        cy.fixture('tripInformations').then((tripInformations) => {
            cy.get('#HeaderPage').contains(tripInformations.endCity)
        })

        //check different activities button
        cy.get('#Hôtels').click()
        cy.wait(500)
        cy.get('#LocationVoitures').click()
        cy.wait(500)
        cy.get('#Activités').click()
        cy.wait(500)
        cy.get('#Restaurants').click()
        cy.wait(500)
        cy.get('#Croisières').click()
        cy.wait(500)

        //Add an activity for each type of activity


        //Delete one activities


        //go to the next page
        cy.get('SaveButton').click()
    })






})