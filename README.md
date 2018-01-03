Basic Introduction to Flogo
===========================     


Create a new application and then a new flow inside the app

1. Select the trigger - install __Timer2__ for the best timer trigger by copying and pasting his Github link `https://github.com/jvanderl/flogo-components/trigger/timer2`

  * If you get any errors when installing a trigger or activity, double check the URL doesn’t include `/tree/master/…`

2. Change the timer settings to be either repeating or non-repeating, start immediately and type in a number in either the hours, minutes or seconds field

3. Add further activities either preinstalled or create your own: To create an activity, follow the steps [here](https://tibcosoftware.github.io/flogo/extensions-developer/create-activity/)
    * To use GPIO pins on the Raspberry Pi, you need:
      *  A __Control GPIO__ activity which sets the method as ‘Direction’, pin number, and direction as either ‘Input’ or ‘Output’
      * Another __Control GPIO__ activity with the method as 'Set State' with the same pin number, no direction required, and state set to ‘High’ or ‘Low’

5. To parse parameters from one activity to another, hover over the activity which requires the parsed value and click on __Transform__. On the field you want changed, click on the box and select the relevant activity’s ‘result’ variable, then save it.
  * To add the parameter into other text, you will need to install a __Combine__ activity (`https://github.com/jvanderl/flogo-components/activity/combine`) and transform the parameter into “part1”.
6. To add a conditional statement, hover over an activity and click on __Add Branch__. Click on the branch to change the condition - the default variable name will be `<activity-name>.result`

7. Once your app is completed and ready to test, return to the list of flows and click on __Build__. To run this on a Raspberry Pi, click on __ARM/Linux__. Save the file and transfer it to your Raspberry Pi: `scp <saved file> <rpi_username>@<rpi_ipaddress>:~/<saved file>`
2. On the Raspberry Pi, make the file executable. Type `chmod +x <saved file>` into terminal

3. Execute the build: `./<saved file>`

### Notes and Common Problems

* Flogo uses the GPIO/BCM number NOT the physical PIN number. See https://pinout.xyz/ for more information.
* Flogo uses __Control GPIO__ activities for digital inputs only - any analog inputs (e.g. temperature) will need a new activity with relevant libraries imported.
* The MQTT inbuilt activity needs all parameters - including the fields labelled ‘optional’.
* When creating your own activity, the package names must be consistent across `activity.go` and `activity.json` otherwise Flogo will error.
* If the activity is added successfully but won’t load when you click on it, copy and paste the JSON headers from a working activity and reinstall.
* __REST__ activities do not work when there are quotes in the URI and Content fields.
* `Runtime error: invalid memory address or nil pointer dereference` - double check in Flogo's WebUI the transformed variables don’t have errors.
* If you want to add more than the default 7 activities in one flow, you will need to export the application as a JSON file and manually add the additional activities. Be aware the activity IDs need to follow on and link IDs need to be added so the activities flow from one to another.
