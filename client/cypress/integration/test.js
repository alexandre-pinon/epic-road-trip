describe('The Register page', () => {





    it('test Twitter', () => {
        cy.visit('https://app2.abtasty.com/login?return_to=%2Fexperiments')
        cy.wait(3000)
        cy.get(':nth-child(2) > [data-testid="inputWrapper"] > .Input__labelTop___30ZNZ').contains('mail')
        cy.get('#email').type('tharick.khudoos@abtasty.com')
        cy.get('#email').should('have.value', 'abtasty')

        cy.get('.FormButtonRow__buttonRow___3k8xJ > .Button__button___1rotk').click()
    })







/*
    it('test Twitter', function () {
        cy.visit('https://www.instagram.com/?hl=fr')
        cy.get('.bIiDR').click()
        cy.get(':nth-child(1) > ._9GP1n > .f0n8F > ._2hvTZ').typ('hehoooo')
        cy.wait(3000)
        cy.get(':nth-child(1) > ._9GP1n > .f0n8F > ._2hvTZ').clear()
        cy.wait(3000)

        //cy.get('#password').type('truc')
        cy.get('.kEKum > :nth-child(3)').should('not.be.disabled')


    })

 */


})