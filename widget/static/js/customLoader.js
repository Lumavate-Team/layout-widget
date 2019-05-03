var customTemplateLoader = {
    loadTemplate: function(name, templateConfig, callback) {
        if (templateConfig.fromUrl) {
            // Uses jQuery's ajax facility to load the markup from a file
            var fullUrl = window.location.origin + templateConfig.fromUrl;
            $.get(fullUrl, function(markupString) {
                // We need an array of DOM nodes, not a string.
                // We can use the default loader to convert to the
                // required format.
                ko.components.defaultLoader.loadTemplate(name, markupString, callback);
            });
        } else {
            // Unrecognized config format. Let another loader handle it.
            callback(null);
        }
    }
  ,loadViewModel: function(name, viewModelConfig, callback) {
        if (viewModelConfig.fromUrl) {
            var fullUrl = window.location.origin + viewModelConfig.fromUrl + '/js';
            $.getScript(fullUrl, function(data, textStatus, jqxhr) {
            // We need a createViewModel function, not a plain constructor.
            // We can use the default loader to convert to the
            // required format.

              ko.components.defaultLoader.loadViewModel(name, new Function(data), callback);
            });
        } else {
            // Unrecognized config format. Let another loader handle it.
            callback(null);
        }
    }
};
 // Register it
ko.components.loaders.unshift(customTemplateLoader);

