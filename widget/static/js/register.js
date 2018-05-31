(function() {
  // Set to promise library to use async validation
  validate.Promise = RSVP.Promise;
  var form = document.getElementById("register");
  var isEditMode = isEditMode();
  form.addEventListener("submit",handleFormSubmit);

  var constraints = {
    firstName: {
      presence: {message: "^Please enter first name."}
    },
    lastName: {
      presence: {message: "^Please enter last name."}
    },
    company: {
      presence: {message: "^Please enter company."}
    },
    emailAddress: {
      presence: {message: "^Please enter email."},
      email: true
    },
    mobileNumber: {
      phoneValidator: true,
    },
  };

  validate.validators.phoneValidator = function (value) {
    return new validate.Promise(function(resolve, reject) {
      // must use resolve with message instead of reject for all validations to run. otherwise you wont get dictionary of errors, just the reject message

      var notifications = document.getElementsByName('sendMobileCheck')[0];
      if (typeof(value) === "undefined" || value === null){
        if(notifications.checked){
          resolve("can't be blank");
        }else{
          resolve();
        }
        return;
      }

      var url = window.location.href + '/validate-phone/' + encodeURIComponent(value);

      httpGetAsync(url, function(status, response){
        if (status == 200){
          var formattedPhoneInput = document.getElementsByName('mobileNumber')[0];
          formattedPhoneInput.value= response.payload.data.formattedPhone;
          resolve();
        } else {
          resolve("is not valid");
        }
      });
    });
  };


  function updateValue(e, elemPrefix){
    var input = document.getElementsByName(elemPrefix + "Check");
    var hidden = document.getElementsByName(elemPrefix + "Notifications");

    hidden[0].value = input[0].checked ? "on" : "off";
  }

  function handleFormSubmit(ev) {
    ev.preventDefault();
    submitData();

    // document.getElementById("submitRegister").disabled = true;
    // var formData = validate.collectFormValues(form);
    // var errors = validate.async(formData, constraints).then(
    //   function(attributes){
    //     submitData();

    //   }, function (errors){
    //     document.getElementById("submitRegister").disabled = false;
    //     if (errors instanceof Error){
    //       console.err("An error occurred", errors);
    //     } else {
    //       showErrors(form, errors);
    //     }
    //   }
    // );
  }

  function showErrors(form, errors){
    errors = errors || {};
    var inputs = document.getElementsByTagName("input");
    for(var i=0,len = inputs.length; i<len; i++){

      showErrorsForInput(inputs[i], errors && errors[inputs[i].name]);
    }
  }

  function showErrorsForInput(input, errors){
    var inputFieldGroup = input.parentNode;

    var messages = inputFieldGroup.getElementsByClassName("messages")
    resetInputMessage(inputFieldGroup);
    if (errors && messages.length > 0) {
      inputFieldGroup.classList.add("has-error");
      for(var i=0,len=errors.length;i<len;i++){
        addError(messages[0], errors[i]);
      }
    }
  }

  function clearErrors(){
    var inputs = form.getElementsByTagName("input");

    for(var i=0,len = inputs.length; i<len; i++){
      resetInputMessage(inputs[i].parentNode);
    }
  }

  function resetInputMessage(inputGroup){
    inputGroup.classList.remove("has-error");
    var errors = inputGroup.getElementsByClassName("error-block");
    for(var i=0,len=errors.length;i<len;i++){
      errors[i].innerText = "";
    }

  }
  function addError(messages, error) {
    messages.classList.add("error-block");
    messages.classList.add("error");
    messages.innerText = error;
  }

  function isEditMode(){
    var imsCookie = getCookie("ims_registered");
    return typeof(imsCookie) !== undefined && imsCookie != null && imsCookie.length > 0;
  }

  function submitData(){
    var message = isEditMode ? "Settings updated!" : "Thank you for registering!";

    clearErrors();
    var url = window.location.href + "/form";
    var formData = validate.collectFormValues(form);
    httpPostAsync(url, JSON.stringify(formData), function(status, response){
      if (status === 204){
        toastr.options.positionClass="toast-top-full-width";
        toastr.success(message);
        window.setTimeout(function(){
          window.location.href= "/";
        },1000);
      }else{
        toastr.error('Sorry, an error occurred');
        console.log(response);
        document.getElementById("submitRegister").disabled = false;
      }
    });
  }

  function httpGetAsync(url, callback)
  {
    var token = getCookie("pwa_jwt");
    var xhr = new XMLHttpRequest();
    xhr.open("GET",url, true);
    xhr.setRequestHeader("Authorization", "Bearer " + token);
    xhr.overrideMimeType("application/json");
    xhr.onreadystatechange = function() { 
        if (xhr.readyState == 4){
            callback(xhr.status, JSON.parse(xhr.responseText));
        }
    }
    xhr.send(null);
  }


  function httpPostAsync(url, formData, callback)
  {
    var token = getCookie("pwa_jwt");
    var xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Authorization", "Bearer " + token);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function() { 
        if (xhr.readyState == 4){
            callback(xhr.status, xhr.responseText);
        }
    }
    xhr.send(formData);
  }

  function getCookie(name){
    var cookies = document.cookie.split(";");
    for(var i=0,len=cookies.length; i < len; i++){
      var cookie = cookies[i].split("=");
      if(cookie[0].trim() == name){
        return cookie[1].trim();
      }
    }
  }
})();