package com.example.projet_collectif_mobile.ui.screens

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.projet_collectif_mobile.models.SurfSpot

import com.example.projet_collectif_mobile.network.SpotsApi
import kotlinx.coroutines.launch

class SpotsViewModel : ViewModel() {/** The mutable State that stores the status of the most recent request */
// Using LiveData to enable observation from the Activity
private val _spotsUiState = MutableLiveData<List<SurfSpot>>()
    val spotsUiState: LiveData<List<SurfSpot>> = _spotsUiState

    init {
        getSpotsList()
    }


    fun getSpotsList() {
        viewModelScope.launch {

            try {
                println("try success")
                val listResult = SpotsApi.retrofitService.getSpots()
                _spotsUiState.value = listResult
                println("reaching Api datas")
            }
            catch (e: Exception) {
                println("error of Api connection")
                println(e)

                _spotsUiState.value = emptyList()
            }
        }
    }
}