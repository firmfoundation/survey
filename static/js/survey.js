function menuSurvey() {
    $(".content").empty();
    $(".content").append(`
        <div class="container-sm">
            <div id="left-col" class="cs-1 was-validated justify-content-center">
                <div class="form-floating w-50">
                    <input type="text" class="form-control form-control-sm" id="txt-survey" value="" size="10" required>
                    <label for="txt-survey" class="form-label">Survey title</label>
                    <div class="invalid-feedback">Survey title or name is required. </div>
                </div>
            
                <div class="mt-3">
                    <button type="button" class="btn btn-warning btn-sm" id="btn-survey-create" onClick="CreateSurvey()">Create Survey</button>
                </div>            
            </div>
        </div>
        <div class="container-sm ">
            <div id="right-col" class="cs-2"></div>
        </div>
    `)
}

function CreateSurvey() {
    var payload = {
        name: $("#txt-survey").val(),
    };

    $.ajax({
        url: '/surveys',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(payload),
        success: function (response) {
           alert("Survey created successfully!")
        },
        error: function (response) {
            alert("error")
        }
    })
}



