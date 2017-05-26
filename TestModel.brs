'//////////////////
'/// Advert Model
'//////////////////
function GetAdvertModel () as Object

    if m._advertModelSingleton = Invalid

        m._advertModelSingleton = UKEventDispatcher()
        m._advertModelSingleton.append({

            '//////////////////
            '/// PUBLIC API ///
            '//////////////////

            ADVERT_CLICKED: "AdvertModel.advertClicked"
            ADVERT_RETRIEVED: "AdvertModel.advertRetrieved"
            FETCH_ADVERT: "AdvertModel.fetchAdvert"
            LOAD_CHANNEL_APP: "AdvertModel.loadChannelApp"


            ' Request advert
            '
            ' @param {String} url
            '
            ' @event    FETCH_ADVERT
            '
            requestAdvert: function (url as String) as Void
                m.dispatchEvent(m.FETCH_ADVERT, m._getTargetedAdvertUrl(url))
            end function


            ' Advert selected (clicked)
            '
            ' @event    ADVERT_CLICKED
            ' @event    LOAD_CHANNEL_APP
            '
            advertSelected: function () as Void
                if m._advert <> Invalid
                    provider = m._providers[m._advert.getPublisherId()]
                    if provider <> Invalid and m._advert.getContentID() <> ""
                        deeplink = DeeplinkData({
                            contentId: m._advert.getContentID()
                            mediaType: m._advert.getMediaType()
                        })

                        if provider.isEnabled()
                            deeplink.setProvider(provider)
                            m.dispatchEvent(m.ADVERT_CLICKED, m._advert.getDFPClickURL())
                        end if

                        m.dispatchEvent(m.LOAD_CHANNEL_APP, deeplink)
                    end if
                end if
            end function


            ' Set advert refresh rate
            '
            ' @param {Integer} refreshRate
            '
            setAdvertRefreshRate: function (refreshRate as Integer) as Void
                m._refreshRate = refreshRate
            end function


            ' Set providers
            '
            ' @param {roAssociativeArray} providers
            '
            setProviders: function (providers as Object) as Void
                m._providers = providers
            end function


            ' Set the advert
            '
            ' @param {roAssociativeArray} response - JSON response from service
            '
            ' @event ADVERT_RETRIEVED @param {AdvertData} advert - Advert data VO
            '
            setAdvert: function (response as Object) as Void
                if TypeIsAssociativeArray(response)
                    m._advert = AdvertData(response.data)
                    m._advert.setRefreshRate(m._refreshRate)

                    m.dispatchEvent(m.ADVERT_RETRIEVED, m._advert)
                end if
            end function

            setDmp : function (dmp as String) as void
                m._dmp = dmp
            end function

            '//////////////////////////
            '/// PRIVATE PROPERTIES ///
            '//////////////////////////

            _QUERY_STRING_PPID_KEY: "ppid="

            _advert: Invalid
            _providers: {}
            _refreshRate: 300
            _dmp: ""


            '///////////////////////
            '/// PRIVATE METHODS ///
            '///////////////////////

            _destroy: function () as Void
                m._clearListeners()
                m._advert = Invalid
                m._providers = {}
            end function


            _getTargetedAdvertUrl: function (url as String) as String
                if Instr(0, url, "?") = 0
                    divider = ";"
                else
                    divider = "&"
                end if

                if m._dmp <> ""
                    encodedDMP = m._dmp.EncodeUriComponent()
                    url = url + divider + m._QUERY_STRING_PPID_KEY + encodedDMP
                end if

                return url
            end function
        })

    end if

    return m._advertModelSingleton

end function


function DestroyAdvertModel () as Void
    if m._advertModelSingleton <> Invalid
        m._advertModelSingleton._destroy()
        m._advertModelSingleton = Invalid
    end if
end function
